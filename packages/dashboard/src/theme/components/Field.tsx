import {
  ChangeEventHandler,
  FC,
  FocusEventHandler,
  PropsWithChildren,
  ReactNode,
  createContext,
  useContext,
  useEffect,
  useState,
} from 'react';
import { Input, InputProps } from './Input';
import styles from './Field.module.scss';
import { Spinner } from './Spinner';
import { Button } from './Button/index';
import { FieldGroupContext } from './FieldGroupContext';
import Chevron from './Chevron.svg?icon';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export type FieldProps<
  T extends
    | 'text'
    | 'date'
    | 'email'
    | 'password'
    | 'color'
    | 'checkbox'
    | 'number' = 'text'
> = Omit<InputProps, 'value' | 'onChange'> & {
  name?: string;
  loading?: boolean;
  type?: T;
  value: T extends 'text' | 'date' | 'email' | 'password' | 'color'
    ? string
    : T extends 'boolean'
    ? boolean
    : T extends 'number'
    ? number
    : never;
  onChange?: T extends 'text' | 'date' | 'email' | 'password' | 'color'
    ? (value: string) => void
    : T extends 'boolean'
    ? (value: boolean) => void
    : T extends 'number'
    ? (value: number) => void
    : never;
  display?: T extends 'text' | 'date' | 'email' | 'password' | 'color'
    ? (value: string) => string
    : T extends 'number'
    ? (value: number) => string
    : never;
} & (T extends 'text' | 'date' | 'email' | 'password' | 'color'
    ? {
        minValue?: never;
        maxValue?: never;
      }
    : T extends 'checkbox'
    ? {
        minValue?: never;
        maxValue?: never;
      }
    : T extends 'number'
    ? {
        minValue?: number;
        maxValue?: number;
      }
    : never);

export const Field: FC<FieldProps<any>> & { Group: FC<FieldGroupProps> } = ({
  type = 'text',
  name,
  loading = false,
  value,
  display,
  leading,
  trailing,
  minValue,
  maxValue,
  onChange,
  ...props
}) => {
  const [_value, setValue] = useState(value);
  const [changed, setChanged] = useState(false);
  const [focused, setFocused] = useState(false);
  const [spin, setSpin] = useState<() => void | undefined>();

  useEffect(() => {
    if (!focused && !changed) {
      setValue(value);
    }
    if (value === _value) {
      setChanged(false);
    }
  }, [value]);

  const adjustValue = (v: string): string => {
    if (!focused) {
      if (type === 'number' && typeof minValue !== 'undefined') {
        if (Number(v) < minValue) {
          v = minValue.toString();
        }
      }

      if (type === 'number' && typeof maxValue !== 'undefined') {
        if (Number(v) > maxValue) {
          v = maxValue.toString();
        }
      }
    }
    return v;
  };

  const onFocus = () => {
    if (type !== 'checkbox') {
      setFocused(true);
    }
  };

  const onBlur = () => {
    setFocused(false);

    if (_value !== value) {
      setChanged(true);
      if (type === 'number') {
        const v = adjustValue(_value as any);
        (onChange as any)?.(v);
      } else {
        (onChange as any)?.(_value);
      }
    } else {
      setChanged(false);
    }
  };

  const _onChange: ChangeEventHandler<HTMLInputElement> = event => {
    if (type === 'checkbox') {
      setValue(!!event.target.checked);

      if (!focused) {
        (onChange as any)?.(!!event.target.checked);
      }
    } else {
      let v = event.target.value;
      v = adjustValue(v);
      setValue(v);

      if (!focused) {
        (onChange as any)?.(v);
      }
    }
  };

  const isInFieldGroup = useContext(FieldGroupContext);

  const onSpinStart = (sign: number) => {
    const handler = () => {
      let value: any;
      setValue(_value => {
        let v = Number(_value);

        const pow10 = Math.pow(10, Math.floor(Math.log10(v)));
        v = (Math.floor((2 * v) / pow10) * pow10) / 2 + (sign * pow10) / 2;
        if (v < (minValue ?? Number.MIN_SAFE_INTEGER)) {
          v = minValue ?? Number.MIN_SAFE_INTEGER;
        }
        if (v > (maxValue ?? Number.MAX_SAFE_INTEGER)) {
          v = maxValue ?? Number.MAX_SAFE_INTEGER;
        }
        return (value = v);
      });
      setImmediate(() => (onChange as any)?.(value));
    };

    let interval: NodeJS.Timer;
    let timeout = setTimeout(() => {
      handler();
      interval = setInterval(handler, 150);
    }, 700);
    setSpin(() => () => {
      clearTimeout(timeout);
      clearInterval(interval);
      handler();
    });
  };

  const onSpinStop = () => {
    spin?.();
    setSpin(undefined);
  };

  return (
    <>
      <Input
        type={type === 'number' ? 'text' : type}
        className={cx('field')}
        {...(type === 'checkbox'
          ? { value: !!_value }
          : {
              value:
                (focused ? _value : (display as any)?.(_value) ?? _value) ?? '',
            })}
        disabled={!name}
        onChange={_onChange}
        onFocus={onFocus}
        onBlur={onBlur}
        leading={
          <>
            {name && (
              <label className={cx({ 'in-field-group': isInFieldGroup })}>
                {name}
              </label>
            )}
            {leading}
          </>
        }
        trailing={
          <>
            <Spinner show={loading} align="right" />
            {type === 'number' && (
              <div className={cx('spinner')}>
                <Button
                  secondary
                  onMouseDown={() => onSpinStart(1)}
                  onMouseUp={onSpinStop}
                >
                  <Chevron width="0.9em" className={cx('chevron')} />
                </Button>
                <Button
                  secondary
                  onMouseDown={() => onSpinStart(-1)}
                  onMouseUp={onSpinStop}
                >
                  <Chevron width="0.9em" className={cx('chevron')} />
                </Button>
              </div>
            )}
            {trailing}
          </>
        }
        {...props}
      />
    </>
  );
};

export interface FieldGroupProps {}

export const FieldGroup: FC<PropsWithChildren<FieldGroupProps>> = ({
  children,
}) => {
  return (
    <div className={cx('field-group')}>
      <FieldGroupContext.Provider value={true}>
        {children}
      </FieldGroupContext.Provider>
    </div>
  );
};

Field.Group = FieldGroup;
