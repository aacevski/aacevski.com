import {

  Button, chakra, Heading, Icon, Link, Stack, Text, useColorModeValue as mode, VStack,
} from '@chakra-ui/react';

import Avatar from '~components/avatar';
import LinkPreview from '~components/link-preview';
import socialLinks from '~constants/social-icons';

const Hero = () => {
  return (
    <VStack spacing={8} fontFamily="JetBrains Mono">
      <VStack spacing={4}>
        <Avatar />
        <Heading fontFamily="inherit" alignSelf="start">Andrej Acevski</Heading>
        <Heading fontFamily="inherit" size="sm" alignSelf="start">
          software engineer, open source advocate, content creator
        </Heading>
        <Text fontFamily="inherit" color="paragraph" fontWeight="normal">
          During the day I&apos;m a software engineer
          {' '}
          <LinkPreview src="/assets/previews/codechem.png">
            <chakra.a
              href="https://codechem.com"
              target="_blank"
              textDecoration="underline"
              textDecorationStyle="wavy"
              textDecorationColor="green.500"
              color={mode('black', 'white')}
              _hover={{
                textDecorationColor: mode('green.700', 'green.300'),
              }}
            >
              @CodeChem
            </chakra.a>
          </LinkPreview>

          .
          {' '}
          At night I&apos;m an instructor
          {' '}
          <LinkPreview src="/assets/previews/egghead.png">
            <chakra.a
              href="https://egghead.io"
              target="_blank"
              textDecoration="underline"
              textDecorationStyle="wavy"
              textDecorationColor="purple.500"
              color={mode('black', 'white')}
              _hover={{
                textDecorationColor: mode('purple.700', 'purple.300'),
              }}
            >
              @Egghead
            </chakra.a>
          </LinkPreview>
          {' '}
          and an collaborator
          {' '}
          <LinkPreview src="/assets/previews/chakra-ui.png">
            <chakra.a
              href="https://chakra-ui.com"
              target="_blank"
              textDecoration="underline"
              textDecorationStyle="wavy"
              textDecorationColor="teal.500"
              color={mode('black', 'white')}
              _hover={{
                textDecorationColor: mode('teal.700', 'teal.300'),
              }}
            >
              @Chakra UI
            </chakra.a>
          </LinkPreview>
          .
        </Text>
      </VStack>
      <Stack
        w="full"
        spacing={4}
        direction={{
          base: 'column',
          md: 'row',
        }}
      >
        {socialLinks.map((socialLink) => (
          <Button
            as={Link}
            key={socialLink.href}
            href={socialLink.href}
            color={socialLink.color}
            target="_blank"
            variant="ghost"
            leftIcon={<Icon fontSize="xl" as={socialLink.icon} />}
          >
            {socialLink.name}
          </Button>
        ))}
      </Stack>
    </VStack>
  );
};

export default Hero;
