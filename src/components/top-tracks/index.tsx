import useSWR from 'swr';
import { Spinner } from '@chakra-ui/react';

import fetcher from '../../../src/utils/fetcher';
import { TrackCard } from '../track-card';
import { Track } from 'src/types/track';

export const TopTracks = () => {
  const { data, error } = useSWR('/api/top-tracks', fetcher);

  if (!data && !error) {
    return <Spinner size="lg" />;
  }

  return data.tracks.map((track: Track, index: number) => (
    <TrackCard
      ranking={index + 1}
      key={track.title}
      image={track.image}
      title={track.title}
      artist={track.artist}
    />
  ));
};
