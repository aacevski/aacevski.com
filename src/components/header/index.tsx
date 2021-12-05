import React from 'react';
import NextLink from 'next/link';
import {
  HStack,
  Link,
  Heading,
  IconButton,
  useColorMode,
  useColorModeValue as mode,
} from '@chakra-ui/react';
import { IoMoon, IoSunny } from 'react-icons/io5';

export const Header = () => {
  const { colorMode, toggleColorMode } = useColorMode();

  return (
    <HStack
      as="nav"
      zIndex="100"
      bg={mode('rgba(255, 255, 255, 0.4)', '#000')}
      position="sticky"
      top="0"
      py={3}
      justifyContent="center"
      w="full"
      style={{
        backdropFilter: 'blur(15px)',
      }}
    >
      <HStack w={{ base: '90%', lg: '50%' }} justifyContent="space-between">
        <HStack spacing={10}>
          <NextLink href="/" passHref>
            <Link>
              <Heading size="xs">Home</Heading>
            </Link>
          </NextLink>
          <NextLink href="/about" passHref>
            <Link>
              <Heading size="xs">About</Heading>
            </Link>
          </NextLink>
          <NextLink href="/bookmarks" passHref>
            <Link>
              <Heading size="xs">Bookmarks</Heading>
            </Link>
          </NextLink>
        </HStack>
        <IconButton
          icon={colorMode === 'light' ? <IoMoon /> : <IoSunny />}
          aria-label="Color mode switch"
          variant="ghost"
          size="md"
          onClick={toggleColorMode}
        />
      </HStack>
    </HStack>
  );
};
