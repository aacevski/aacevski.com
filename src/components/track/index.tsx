import { Flex, Text, Image, useColorModeValue as mode } from '@chakra-ui/react';

const Track = (track: any) => {
  return (
    <Flex w="full" direction="column" mt={4}>
      <Flex
        alignItems="center"
        mb={4}
        _hover={{ bg: mode('blackAlpha.100', 'blackAlpha.500') }}
        py={4}
      >
        <Flex alignItems="center" w={10} mr={4}>
          <Text fontSize="sm" color="#7e7e7e" fontWeight="bold" ml={4}>
            {track.ranking}
          </Text>
        </Flex>
        <Image
          src={track.image}
          w="50px"
          objectFit="cover"
          h="64px"
          rounded="md"
          mr={4}
          alt={track.title}
        />
        <Flex direction="column">
          <Text fontSize="lg" fontWeight="bold">
            {track.title}
          </Text>
          <Text>{track.artist}</Text>
        </Flex>
      </Flex>
    </Flex>
  );
};

export default Track;
