import { createElement, ElementType, forwardRef, useContext } from 'react';
import {
  useSelect,
  SelectProps as SelectUnstyledProps,
  SelectOptionDefinition,
} from '@mui/base';
import styles from './Select.module.scss';
import { FieldGroupContext } from './FieldGroupContext';
import Chevron from './Chevron.svg?icon';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export type SelectProps<
  K extends {} = any,
  T extends ElementType = 'select'
> = SelectUnstyledProps<K, false, T> & {
  component?: ElementType;
  control?: boolean;
  options: SelectOptionDefinition<K>[];
};

export const Select = forwardRef<HTMLSelectElement, SelectProps>(
  (
    {
      component = 'select',
      className,
      control = true,
      value,
      onChange,
      placeholder,
      options,
      ...props
    },
    ref
  ) => {
    const { disabled, getButtonProps, getListboxProps } = useSelect({
      options,
      ...props,
    });

    const isInFieldGroup = useContext(FieldGroupContext);

    return (
      <div
        className={cx(className, 'select', {
          disabled,
          control,
        })}
      >
        <div className={cx('wrapper', { 'in-field-group': isInFieldGroup })}>
          {createElement(
            component,
            {
              ...props,
              ...getButtonProps(),
              // ...getListboxProps(),
              value,
              onChange,
              placeholder,
            },
            options.map(option => (
              <option
                key={option.value}
                value={option.value}
                disabled={option.disabled}
              >
                {option.label}
              </option>
            ))
          )}
          <Chevron width="1em" className={cx('chevron')} />
        </div>
      </div>
    );
  }
);
