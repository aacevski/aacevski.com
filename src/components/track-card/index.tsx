import {
  HStack,
  Text,
  useColorModeValue as mode,
  VStack,
} from '@chakra-ui/react';

import Image from 'src/components/image';
import { Track } from 'src/types/track';

export const TrackCard = ({ ranking, image, title, artist }: Track) => {
  return (
    <HStack
      w="full"
      spacing={6}
      py={4}
      _hover={{ bg: mode('blackAlpha.100', 'blackAlpha.500') }}
    >
      <Text fontSize="sm" color="#7e7e7e" fontWeight="bold" ml={4}>
        {ranking}
      </Text>
      <Image
        src={image}
        w="50px"
        objectFit="cover"
        h="64px"
        rounded="md"
        mr={4}
        alt={title}
      />
      <VStack align="flex-start" spacing={0}>
        <Text fontSize="lg" fontWeight="bold">
          {title}
        </Text>
        <Text>{artist}</Text>
      </VStack>
    </HStack>
  );
};
