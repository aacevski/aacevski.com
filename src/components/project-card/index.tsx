import React from 'react';
import {
  VStack,
  Heading,
  LinkBox,
  LinkOverlay,
  AspectRatio,
} from '@chakra-ui/react';

import { Project } from '../../types/project';
import Image from 'src/components/image';

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
          <AspectRatio ratio={16 / 9} w="full">
            <Image src={preview} alt={title} />
          </AspectRatio>
        </LinkOverlay>
      </VStack>
    </LinkBox>
  );
};
