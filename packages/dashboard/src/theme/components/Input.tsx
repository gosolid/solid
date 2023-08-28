import {
  createElement,
  ElementType,
  forwardRef,
  PropsWithChildren,
  ReactElement,
  useContext,
} from 'react';
import { useInput, InputProps as InputUnstyledProps } from '@mui/base';
import styles from './Input.module.scss';
import { FieldGroupContext } from './FieldGroupContext';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export type InputProps<T extends ElementType = 'input'> =
  InputUnstyledProps<T> & {
    leading?: ReactElement;
    trailing?: ReactElement;
    control?: boolean;
    multiple?: boolean;
  };

export const Input = forwardRef<HTMLInputElement, InputProps<any>>(
  (
    {
      component = 'input',
      leading,
      trailing,
      className,
      control = true,
      type,
      autoComplete,
      spellCheck,
      autoCapitalize,
      value,
      multiple,
      onChange,
      placeholder,
      ...props
    },
    ref
  ) => {
    const { disabled, required, focused, error, getRootProps, getInputProps } =
      useInput(props);

    const isInFieldGroup = useContext(FieldGroupContext);

    const valueProp: any = {};
    if (type === 'checkbox') {
      valueProp.checked = value;
    } else {
      valueProp.value = value;
    }

    return (
      <div
        className={cx(className, 'input', {
          disabled,
          required,
          focused,
          error,
          control,
        })}
      >
        {leading}
        <div className={cx('wrapper', { 'in-field-group': isInFieldGroup })}>
          {createElement(component, {
            ...getRootProps(),
            ...getInputProps(),
            type,
            multiple,
            autoComplete,
            spellCheck,
            autoCapitalize,
            ...valueProp,
            onChange,
            placeholder,
          })}
        </div>
        {trailing}
      </div>
    );
  }
);
