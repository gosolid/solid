import { ComponentType, FC, ReactNode, useEffect, useState } from 'react';
import { Photo } from './Photo';
import _RenderIfVisible from 'react-render-if-visible';
import format from 'date-fns/format';
import styles from './BlogIndex.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

const { default: RenderIfVisible } = _RenderIfVisible as any;

export interface ApplePhotoDerivative {
  checksum: string;
  url: string;
}

export interface ApplePhoto {
  photoGuid: string;
  dateCreated: string;
  derivatives: Record<string, ApplePhotoDerivative>;
}

export type AppleWebStream = {
  photos: ApplePhoto[];
};

export const fetchApplePhotosAlbum = async (albumId: string) => {
  let response;

  response = await fetch(
    `https://apple.photos.grexie.com/${albumId}/sharedstreams/webstream`,
    {
      method: 'POST',
      headers: {
        'Content-Type': 'text/plain',
        Accept: '*/*',
      },
      credentials: 'include',
      body: JSON.stringify({
        streamCtag: null,
      }),
    }
  );
  const webstream: AppleWebStream = await response.json();

  const photoGuids = webstream.photos.map(photo => photo.photoGuid);

  response = await fetch(
    `https://apple.photos.grexie.com/${albumId}/sharedstreams/webasseturls`,
    {
      method: 'POST',
      headers: {
        'Content-Type': 'text/plain',
        Accept: '*/*',
      },
      credentials: 'include',
      body: JSON.stringify({ photoGuids }),
    }
  );

  const assets = await response.json();

  webstream.photos.forEach(photo => {
    for (const k in photo.derivatives) {
      const { checksum } = photo.derivatives[k];
      const asset = assets.items[checksum];
      photo.derivatives[
        k
      ].url = `https://${asset.url_location}${asset.url_path}`;
    }
  });

  return webstream;
};

export const fetchApplePhotos = async (sharedAlbums: string[]) => {
  return Promise.all(
    sharedAlbums.map(album => fetchApplePhotosAlbum(album.replace(/^.*#/, '')))
  );
};

export interface BlogIndexProps {
  albums: string[];
}

const renderPhoto = (photo: ApplePhoto) => {
  const el = (
    <li key={photo.photoGuid} className={cx('photo')}>
      <RenderIfVisible>
        <Photo photo={photo} />
      </RenderIfVisible>
    </li>
  );
  return { el, date: new Date(photo.dateCreated) };
};

type Post = { date: Date; el: ReactNode };

const addDays = (posts: Post[]) => {
  let lastDate;
  const out = [];

  for (const post of posts) {
    if (
      !lastDate ||
      post.date.getFullYear() != lastDate.getFullYear() ||
      post.date.getMonth() != lastDate.getMonth() ||
      post.date.getDay() != lastDate.getDay()
    ) {
      out.push({
        date: post.date,
        el: (
          <li key={`date-${post.date.toISOString()}`}>
            <h2>{`${format(post.date, 'E do MMMM yyyy')}`}</h2>
          </li>
        ),
      });
      lastDate = post.date;
    }
    out.push(post);
  }

  return out;
};

export const BlogIndex: FC<BlogIndexProps> = ({ albums: _albums }) => {
  const [albums, setAlbums] = useState<AppleWebStream[]>([]);

  useEffect(() => {
    if (!_albums?.length) {
      return;
    }

    fetchApplePhotos(_albums)
      .then(albums => setAlbums(albums))
      .catch(err => console.error(err));
  }, []);

  let photos = albums.reduce(
    (a, album) => [...a, ...album.photos.map(renderPhoto)],
    [] as Post[]
  );

  const posts = addDays(
    [...photos].sort((a, b) => b.date.getTime() - a.date.getTime())
  ).map(a => a.el);

  console.info(posts);
  return <ul className={cx('container')}>{posts}</ul>;
};
