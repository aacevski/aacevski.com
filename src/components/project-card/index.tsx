import React from 'react';
import { VStack, Heading, LinkBox, LinkOverlay } from '@chakra-ui/react';

import Image from 'src/components/image';
import { Project } from 'src/types/project';

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
          <Image
            width={1920}
            height={1080}
            src={preview}
            alt={title}
            priority
            rounded="lg"
          />
        </LinkOverlay>
      </VStack>
    </LinkBox>
  );
};
