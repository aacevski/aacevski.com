import {
  HStack,
  Text,
  useColorModeValue as mode,
  VStack,
  chakra,
} from '@chakra-ui/react';

import Image from 'src/components/image';
import { Track } from 'src/types/track';

export const TrackCard = ({ ranking, image, title, artist }: Track) => {
  return (
    <HStack
      w="full"
      spacing={{ base: 6, md: 6 }}
      py={4}
      _hover={{ bg: mode('blackAlpha.100', 'blackAlpha.500') }}
    >
      <Text fontSize="sm" color="#7e7e7e" fontWeight="bold">
        {ranking}
      </Text>
      <chakra.span flexShrink={0}>
        <Image
          src={image}
          width={50}
          height={64}
          objectFit="cover"
          rounded="md"
          alt={title}
        />
      </chakra.span>

      <VStack align="flex-start">
        <Text fontSize="lg" fontWeight="bold">
          {title}
        </Text>
        <Text>{artist}</Text>
      </VStack>
    </HStack>
  );
};
