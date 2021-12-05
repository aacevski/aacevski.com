import useSWR from 'swr';
import { Spinner } from '@chakra-ui/react';

import fetcher from '../../../src/utils/fetcher';
import Track from '../track';

export const TopTracks = () => {
  const { data, error } = useSWR('/api/top-tracks', fetcher);

  if (!data && !error) {
    return <Spinner size="lg" />;
  }

  return data.tracks.map((track: any, index: any) => (
    <Track ranking={index + 1} key={track.songUrl} {...track} />
  ));
};
