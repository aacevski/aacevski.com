import React from 'react';
import { VStack, Heading, Box, LinkBox, LinkOverlay } from '@chakra-ui/react';

import { Project } from '../../types/project';

export const ProjectCard = ({ title, description, preview, url }: Project) => {
  return (
    <LinkBox>
      <VStack align="flex-start">
        <Heading mb={2}>{title}</Heading>
        <Heading fontSize="lg" color="gray" fontWeight="400" mb={8}>
          {description}
        </Heading>
        <VStack w="full" justifyContent="center" />
        <LinkOverlay href={url} isExternal>
          <Box
            as="video"
            w="full"
            rounded="3xl"
            playsInline
            autoPlay
            muted
            loop
          >
            <source src={preview} type="video/mp4" />
          </Box>
        </LinkOverlay>
      </VStack>
    </LinkBox>
  );
};
