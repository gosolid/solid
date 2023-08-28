import { useState, useEffect, useMemo, useRef, FC } from 'react';
import { createPortal } from 'react-dom';
import styles from './Photo.module.scss';
import format from 'date-fns/format';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface VideoProps {
  video: any;
}

export const Video: FC<VideoProps> = ({ video }) => {
  const el = useRef<HTMLVideoElement>(null);
  const [show, setShow] = useState(false);
  const fullscreen = useMemo(() => {
    const container = document.createElement('div');
    document.body.appendChild(container);
    return container;
  }, []);
  const elfs = useRef<HTMLVideoElement>(null);

  useEffect(() => {
    if (!el.current) {
      return;
    }

    el.current.addEventListener(
      'mouseenter',
      () => {
        try {
          el.current?.play().catch(() => {});
        } catch (err) {}
      },
      false
    );

    el.current.addEventListener(
      'mouseleave',
      () => {
        try {
          el.current?.pause();
        } catch (err) {}
      },
      false
    );

    el.current.addEventListener(
      'click',
      () => {
        setShow(true);
        try {
          elfs.current?.play().catch(() => {});
          if (el.current) {
            el.current?.pause();
            el.current.currentTime = 0;
          }
        } catch (err) {}
      },
      false
    );

    fullscreen.addEventListener(
      'click',
      () => {
        if (elfs.current) {
          elfs.current?.pause();
          elfs.current.currentTime = 0;
          elfs.current.load();
        }
        setShow(false);
      },
      false
    );

    elfs.current?.addEventListener('pause', () => {
      if (elfs.current) {
        elfs.current?.pause();
        elfs.current.currentTime = 0;
      }
      setShow(false);
    });

    return () => {
      fullscreen.remove();
    };
  }, []);

  return (
    <>
      <video
        poster={video.derivatives.PosterFrame.url}
        muted
        loop
        playsInline
        ref={el}
      >
        <source src={video.derivatives['360p'].url} />
      </video>
      {createPortal(
        <div className={cx('fullscreen', { show })}>
          <video
            poster={video.derivatives.PosterFrame.url}
            loop
            playsInline
            ref={elfs}
          >
            <source src={video.derivatives['720p'].url} />
          </video>
          <div className={cx('meta')}>
            <div className={cx('meta-author')}>
              {!!video.contributorFirstName && (
                <div className={cx('name')}>{video.contributorFirstName}</div>
              )}
              {!!video.dateCreated && (
                <div className={cx('date')}>
                  {format(new Date(video.dateCreated), 'E do MMMM yyyy')}
                </div>
              )}
            </div>
            {!!video.caption && (
              <div className={cx('caption')}>{video.caption}</div>
            )}
          </div>
        </div>,
        fullscreen
      )}
    </>
  );
};

export interface PhotoProps {
  photo: any;
}

export const Photo: FC<PhotoProps> = ({ photo }) => {
  if (photo.mediaAssetType === 'video') {
    return <Video video={photo} />;
  }

  const el = useRef<HTMLImageElement>(null);
  const [show, setShow] = useState(false);
  const fullscreen = useMemo(() => {
    const container = document.createElement('div');
    document.body.appendChild(container);
    return container;
  }, []);

  useEffect(() => {
    el.current?.addEventListener(
      'click',
      () => {
        setShow(true);
      },
      false
    );

    fullscreen.addEventListener(
      'click',
      () => {
        setShow(false);
      },
      false
    );
  }, []);

  const derivative = photo.derivatives[Object.keys(photo.derivatives)[0]];
  const lastDerivative =
    photo.derivatives[Object.keys(photo.derivatives).reverse()[0]];

  return (
    <>
      <img src={derivative.url} ref={el} />
      {createPortal(
        <div className={cx('fullscreen', { show })}>
          <img src={lastDerivative.url} />
          <div className={cx('meta')}>
            <div className={cx('meta-author')}>
              {!!photo.contributorFirstName && (
                <div className={cx('name')}>{photo.contributorFirstName}</div>
              )}
              {!!photo.dateCreated && (
                <div className={cx('date')}>
                  {format(new Date(photo.dateCreated), 'E do MMMM yyyy')}
                </div>
              )}
            </div>
            {!!photo.caption && (
              <div className={cx('caption')}>{photo.caption}</div>
            )}
          </div>
        </div>,
        fullscreen
      )}
    </>
  );
};
