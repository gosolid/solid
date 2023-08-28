import { FC } from 'react';
import styles from './YouTube.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface YouTubeProps {
  url: string;
  orientation?: 'portrait' | 'landscape' | 'square';
}

export const YouTube: FC<YouTubeProps> = ({
  url,
  orientation = 'portrait',
}) => {
  return (
    <div className={cx('youtube', `youtube-${orientation}`)}>
      <div>
        <iframe
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
          allowFullScreen
          loading="lazy"
          width="100%"
          height="100%"
          src={url.replace(
            /^https:\/\/youtu.be\//,
            'https://www.youtube.com/embed/'
          )}
        />
      </div>
    </div>
  );
};
