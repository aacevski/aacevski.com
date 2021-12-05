import React from 'react';
import { GetStaticProps } from 'next';
import {
  Avatar,
  Box,
  Button,
  Heading,
  Icon,
  Link,
  HStack,
  VStack,
  useColorModeValue as mode,
} from '@chakra-ui/react';
import { EmailIcon } from '@chakra-ui/icons';
import { FaLinkedin } from 'react-icons/fa';
import { GoMarkGithub } from 'react-icons/go';
import { SiHashnode } from 'react-icons/si';

import { ProjectCard } from '../src/components/project-card';
import { Project } from '../src/types/project';
import { readData } from 'src/utils/get-projects';

type Props = {
  projects: Project[];
};

const IndexPage = ({ projects }: Props) => {
  return (
    <VStack w="full" spacing={16}>
      <VStack spacing={8}>
        <VStack spacing={6}>
          <Avatar src="./images/andrej.jpg" alignSelf="flex-start" size="lg" />
          <Heading>
            Hi, I’m Andrej. I’m a{' '}
            <Box
              as="span"
              bgGradient="linear(to-l, #12c2e9, #c471ed, #f64f59)"
              bgClip="text"
              textDecoration="underline"
            >
              web developer
            </Box>
            , located in Macedonia.
          </Heading>
        </VStack>
        <HStack spacing={6} alignSelf="flex-start">
          <Button
            id="gradientButton"
            variant="ghost"
            px={8}
            py={6}
            rounded="xl"
            _hover={{
              transform: 'scale(1.1)',
              transitionDuration: '0.3s',
              textDecoration: 'none',
            }}
            border="solid 1px transparent"
            leftIcon={<EmailIcon />}
            href="mailto:aacevski@gmail.com"
            as={Link}
            boxShadow={mode(
              '2px 1000px 1px white inset',
              '2px 1000px 1px black inset'
            )}
          >
            Contact
          </Button>
          <Link
            href="https://github.com/aacevski"
            _active={{ textDecoration: 'none' }}
            _focus={{ textDecoration: 'none' }}
          >
            <Icon
              as={GoMarkGithub}
              w={5}
              h={5}
              color={mode('black', 'white')}
              transitionDuration="0.3s"
              cursor="pointer"
              _hover={{
                transform: 'translate(0px, -5px)',
              }}
            />
          </Link>
          <Link
            href="https://www.linkedin.com/in/andrej-acevski-6292ba179/"
            _active={{ textDecoration: 'none' }}
            _focus={{ textDecoration: 'none' }}
          >
            <Icon
              as={FaLinkedin}
              w={5}
              h={5}
              color={mode('black', 'white')}
              transitionDuration="0.3s"
              cursor="pointer"
              _hover={{
                transform: 'translate(0px, -5px)',
              }}
            />
          </Link>
          <Link
            href="https://aacevski.hashnode.dev/"
            _active={{ textDecoration: 'none' }}
            _focus={{ textDecoration: 'none' }}
          >
            <Icon
              as={SiHashnode}
              w={5}
              h={5}
              color={mode('black', 'white')}
              transitionDuration="0.3s"
              cursor="pointer"
              _hover={{
                transform: 'translate(0px, -5px)',
              }}
            />
          </Link>
        </HStack>
      </VStack>
      <VStack spacing={12}>
        {projects.map((project: Project) => (
          <ProjectCard
            key={project.title}
            title={project.title}
            description={project.description}
            preview={project.preview}
            url={project.url}
          />
        ))}
      </VStack>
    </VStack>
  );
};

export const getStaticProps: GetStaticProps<Props> = async () => {
  const { projects } = await readData<{
    projects: Project[];
  }>('data/projects.json');

  const props: Props = { projects };

  return { props };
};

export default IndexPage;
