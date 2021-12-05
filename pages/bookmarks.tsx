import { useState } from 'react';
import { GetStaticProps } from 'next';
import {
  Button,
  Icon,
  VStack,
  Text,
  Heading,
  HStack,
  useColorModeValue as mode,
  useBreakpointValue,
} from '@chakra-ui/react';
import { DiReact, DiCss3, DiHtml5 } from 'react-icons/di';
import { BsSearch } from 'react-icons/bs';

import { BookmarkCard } from '../src/components/bookmark-card';
import { Bookmark } from '../src/types/bookmark';

type Props = {
  bookmarks: Bookmark[];
  tags: string[];
};

const Bookmarks = ({ bookmarks, tags }: Props) => {
  const [displayBookmarks, setDisplayBookmarks] = useState(bookmarks);
  const [selectedTag, setSelectedTag] = useState<string>();
  const isMediumResolution = useBreakpointValue({ md: true });
  const filterBookmarks = (tag?: string) => {
    if (tag) {
      setDisplayBookmarks(bookmarks.filter(({ tags }) => tags.includes(tag)));
    } else {
      setDisplayBookmarks(bookmarks);
    }
    setSelectedTag(tag);
  };

  return (
    <VStack w="full" align="flex-start" spacing={10}>
      <VStack align="flex-start">
        <Text fontSize="3rem" mb={2}>
          👾
        </Text>
        <Heading size="lg">Bookmarks</Heading>
        <Heading fontSize="lg" color="gray" fontWeight="400" mb={8}>
          Resources that I always lean on.
        </Heading>
      </VStack>
      <HStack w="full" spacing={{ base: 5, md: 10 }} justify="center">
        <Button
          w={{ base: 'full', md: 28 }}
          id={!selectedTag ? 'gradientButton' : ''}
          variant="ghost"
          py={4}
          rounded="xl"
          _hover={{
            transform: 'scale(1.1)',
            transitionDuration: '0.3s',
          }}
          border="solid 1px transparent"
          boxShadow={mode(
            '2px 1000px 1px white inset',
            '2px 1000px 1px black inset'
          )}
          _focus={{
            bg: 'transparent',
          }}
          leftIcon={
            isMediumResolution && <Icon as={BsSearch} width={4} height={4} />
          }
          onClick={() => {
            filterBookmarks();
          }}
        >
          {isMediumResolution ? (
            'All'
          ) : (
            <Icon as={BsSearch} width={4} height={4} />
          )}
        </Button>
        <Button
          id={selectedTag === 'HTML' ? 'gradientButton' : ''}
          variant="ghost"
          w={{ base: 'full', md: 28 }}
          py={4}
          rounded="xl"
          _focus={{
            bg: 'transparent',
          }}
          _hover={{
            transform: 'scale(1.1)',
            transitionDuration: '0.3s',
          }}
          border="solid 1px transparent"
          boxShadow={mode(
            '2px 1000px 1px white inset',
            '2px 1000px 1px black inset'
          )}
          leftIcon={
            isMediumResolution && <Icon as={DiHtml5} width={4} height={4} />
          }
          onClick={() => {
            filterBookmarks('HTML');
          }}
        >
          {isMediumResolution ? (
            'HTML'
          ) : (
            <Icon as={DiHtml5} width={4} height={4} />
          )}
        </Button>
        <Button
          id={selectedTag === 'CSS' ? 'gradientButton' : ''}
          variant="ghost"
          w={{ base: 'full', md: 28 }}
          py={4}
          rounded="xl"
          _focus={{
            bg: 'transparent',
          }}
          _hover={{
            transform: 'scale(1.1)',
            transitionDuration: '0.3s',
          }}
          border="solid 1px transparent"
          boxShadow={mode(
            '2px 1000px 1px white inset',
            '2px 1000px 1px black inset'
          )}
          leftIcon={
            isMediumResolution && <Icon as={DiCss3} width={4} height={4} />
          }
          onClick={() => {
            filterBookmarks('CSS');
          }}
        >
          {isMediumResolution ? (
            'CSS'
          ) : (
            <Icon as={DiCss3} width={4} height={4} />
          )}
        </Button>
        <Button
          id={selectedTag === 'React' ? 'gradientButton' : ''}
          variant="ghost"
          w={{ base: 'full', md: 28 }}
          py={4}
          rounded="xl"
          _focus={{
            bg: 'transparent',
          }}
          _hover={{
            transform: 'scale(1.1)',
            transitionDuration: '0.3s',
          }}
          border="solid 1px transparent"
          boxShadow={mode(
            '2px 1000px 1px white inset',
            '2px 1000px 1px black inset'
          )}
          leftIcon={
            isMediumResolution && <Icon as={DiReact} width={6} height={6} />
          }
          onClick={() => {
            filterBookmarks('React');
          }}
        >
          {isMediumResolution ? (
            'React'
          ) : (
            <Icon as={DiReact} width={6} height={6} />
          )}
        </Button>
      </HStack>
      <VStack align="center" w="full" spacing={6} justify="center">
        {displayBookmarks.map((bookmark: Bookmark, index: number) => (
          <BookmarkCard
            index={index}
            key={bookmark.title}
            title={bookmark.title}
            excerpt={bookmark.excerpt}
            link={bookmark.link}
            cover={bookmark.cover}
            tags={bookmark.tags}
          />
        ))}
      </VStack>
    </VStack>
  );
};

export const getStaticProps: GetStaticProps<Props> = async () => {
  const req = await fetch(
    `https://api.raindrop.io/rest/v1/raindrops/${process.env.RAINDROP_COLLECTION}`,
    {
      headers: {
        Authorization: `Bearer ${process.env.RAINDROP_TOKEN}`,
      },
    }
  );

  const data = await req.json();

  const bookmarks: Bookmark[] = data.items.map(
    ({ cover, title, link, tags, excerpt }) => ({
      link,
      title,
      cover,
      tags,
      excerpt,
    })
  );

  const tags = Array.from(new Set(bookmarks.flatMap(({ tags }) => tags)));

  const props: Props = { bookmarks, tags };

  return {
    props,
    revalidate: 60 * 60,
  };
};

export default Bookmarks;
