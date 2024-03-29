import { Box, Container, Heading, HStack, Stack, Text, Icon, LinkOverlay, LinkBox } from '@chakra-ui/react';
import matter from 'gray-matter';
import { GetStaticPaths, GetStaticProps } from 'next';
import { MDXRemote, MDXRemoteSerializeResult } from 'next-mdx-remote';
import { serialize } from 'next-mdx-remote/serialize';
import readingTimeParser from 'reading-time';
import { useRouter } from 'next/router';
import { useEffect } from 'react';
import { NextSeo } from 'next-seo';
import NextLink from 'next/link';
import { CgPushLeft } from 'react-icons/cg';

import MDXComponents from '~components/mdx-components';
import { BlogPost } from '~types/blog-post';
import { getBlogPosts } from '~utils/get-blog-posts';
import { readBlogPost } from '~utils/read-blog-post';
import useBlogPostViews from '~hooks/use-blog-post-views';

type Props = BlogPost & {
  source: MDXRemoteSerializeResult;
};

const BlogPostPage = ({
  source,
  title,
  readingTime,
  date,
  description,
}: Props) => {
  const { query } = useRouter();
  const slug = query.slug as string;

  const { views, increment: incrementViews, isLoading: isLoadingViews } = useBlogPostViews(slug);
  const url = `https://aacevski.com/blog/${slug}`;
  const ogImage = `https://aacevski.com/assets/blog/${slug}.png`;

  useEffect(() => {
    if (slug) {
      incrementViews();
    }
  }, [slug, incrementViews]);

  return (
    <>
      <NextSeo
        title={`${title} - Andrje Acevski`}
        description={description}
        canonical={url}
        openGraph={{
          description,
          title: `${title} - Andrej Acevski`,
          url: `https://aacevski.com/blog/${slug}`,
          images: [
            {
              url: ogImage,
              height: 630,
              width: 1200,
              alt: title,
            },
          ],
        }}
      />
      {/* go back button */}
      <Container maxW="container.lg" mb={10}>
        <LinkBox w="max-content">
          <LinkOverlay display="flex" mb={6} alignItems="center" as={NextLink} href="/" gap={2}>
            <Icon as={CgPushLeft} color="paragraph" fontSize="xl" />
            <Text
              fontSize="sm"
              color="paragraph"
              fontFamily="JetBrains Mono"
            >
              /home
            </Text>
          </LinkOverlay>
        </LinkBox>
        {/* </HStack> */}
        <Heading size="xl">
          {title}
        </Heading>
        <Stack
          direction={{ base: 'column', md: 'row' }}
          fontFamily="JetBrains Mono"
          mt={2}
          mb={8}
          divider={<HStack mx={2} />}
          fontSize="sm"
        >
          <Text>
            📅
            {' '}
            {date}
          </Text>
          <Text>
            🕑
            {' '}
            {readingTime}
          </Text>
          {(isLoadingViews || !views) ? (
            <Box display="inline">
              👀
              {' '}
              Loading views...
            </Box>
          ) : (
            <Text>
              👀
              {' '}
              {views}
              {' '}
              views
            </Text>
          )}
        </Stack>
        {/* eslint-disable-next-line @typescript-eslint/no-explicit-any */}
        <MDXRemote lazy {...source} components={MDXComponents as any} />
      </Container>
    </>
  );
};

export const getStaticPaths: GetStaticPaths = async () => {
  const posts = await getBlogPosts();

  return {
    paths: posts.map(({ slug }) => ({ params: { slug } })),
    fallback: false,
  };
};

export const getStaticProps: GetStaticProps<Props> = async (ctx) => {
  const slug = ctx?.params?.slug as string;

  const postContent = await readBlogPost(slug);
  const {
    content,
    data: { title, description, date },
  } = matter(postContent);

  return {
    props: {
      source: await serialize(content),
      readingTime: readingTimeParser(content).text,
      title,
      description,
      date,
      slug,
    },
  };
};

export default BlogPostPage;
