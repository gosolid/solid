import {
  createElement,
  ElementType,
  forwardRef,
  PropsWithChildren,
  FormEvent,
} from 'react';
import {
  useFormControlContext as useFormControlUnstyledContext,
  FormControlProps as FormControlUnstyledProps,
} from '@mui/base';
import styles from './Input.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export type FormProps<T extends ElementType> = FormControlUnstyledProps<T>;

export const Form = forwardRef<any, PropsWithChildren<FormProps<any>>>(
  ({ component = 'form', children, className }, ref) => {
    const props = useFormControlUnstyledContext();

    const onSubmit = (event: FormEvent) => {
      event.preventDefault();
    };

    return createElement(
      component,
      {
        ...props,
        onSubmit,
        className: cx(className, 'form', {
          disabled: props?.disabled,
          required: props?.required,
          focused: props?.focused,
          error: props?.error,
          filled: props?.filled,
        }),
      },
      children
    );
  }
);
