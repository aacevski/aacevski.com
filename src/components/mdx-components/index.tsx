/* eslint-disable no-param-reassign */
/* eslint-disable react/no-array-index-key */
import {
  Alert,
  AlertProps,
  Box,
  BoxProps,
  chakra, ChakraComponent, HTMLChakraProps,
  Kbd, Link, useClipboard, useColorModeValue,
} from '@chakra-ui/react';
import NextImage, { ImageProps } from 'next/image';
import Highlight, { defaultProps, Language } from 'prism-react-renderer';
import darkTheme from 'prism-react-renderer/themes/nightOwl';
import lightTheme from 'prism-react-renderer/themes/nightOwlLight';
import { useId } from 'react';
import slugify from 'slugify';
import { MdContentCopy } from 'react-icons/md';
import { IoMdCheckmark } from 'react-icons/io';
import { MDXRemoteProps } from 'next-mdx-remote';

const ChakraHighlight = chakra(Highlight, {
  shouldForwardProp: (prop) => ['Prism', 'theme', 'code', 'language', 'children'].includes(prop),
});

const Pre = (props: ChakraComponent<'div'>) => <chakra.div my="2em" borderRadius="sm" {...props} />;

const Table = (props: ChakraComponent<'table'>) => (
  <chakra.div overflowX="auto">
    <chakra.table textAlign="left" mt="32px" width="full" {...props} />
  </chakra.div>
);

const THead = (props: ChakraComponent<'th'>) => (
  <chakra.th
    bg="gray.50"
    _dark={{
      bg: 'whiteAlpha.100',
    }}
    fontWeight="semibold"
    p={2}
    fontSize="sm"
    {...props}
  />
);

const TData = (props: ChakraComponent<'td'>) => (
  <chakra.td
    p={2}
    borderTopWidth="1px"
    borderColor="inherit"
    fontSize="sm"
    whiteSpace="normal"
    {...props}
  />
);

const CodeHighlight = ({ children: codeString, className: language }: {
  children: string; className: Language;
}) => {
  const theme = useColorModeValue(lightTheme, darkTheme);
  const codeId = useId();
  const { hasCopied, onCopy } = useClipboard('');

  if (!language) {
    return (
      <chakra.code
        apply="mdx.code"
        color="blue.500"
        _dark={{
          color: 'blue.200',
          bg: 'blue.900',
        }}
        bg="blue.50"
        px={1}
        py={0.5}
        rounded={{ base: 'none', md: 'md' }}
      >
        {codeString}
      </chakra.code>
    );
  }

  language = language.replace('language-', '') as Language;
  const showLineNumbers = !['shell', 'text'].includes(language);

  return (
    <ChakraHighlight
      {...defaultProps}
      code={codeString}
      language={language}
      theme={theme}
    >
      {({ className, style, tokens, getLineProps, getTokenProps }) => {
        tokens.pop();
        return (
          <chakra.div data-language={className} pos="relative" role="group">
            <chakra.pre
              className={className}
              sx={{ ...style, backgroundColor: 'blackAlpha.50' }}
              overflowX="auto"
              rounded="md"
              p={4}
              mx={-4}
              _dark={{
                backgroundColor: 'whiteAlpha.50',
              }}
            >
              {tokens.map((line, i) => {
                const lineProps = getLineProps({ line, key: i });
                return (
                  <chakra.div
                    {...lineProps}
                    display="table-row"
                    key={`${codeId}.${i}`}
                  >
                    {showLineNumbers && (
                      <chakra.span
                        key={`${codeId}.${i}.number`}
                        w={8}
                        display="table-cell"
                        textAlign="right"
                        userSelect="none"
                        color="blackAlpha.500"
                        pr={3}
                        _dark={{
                          color: 'whiteAlpha.500',
                        }}
                      >
                        {i + 1}
                      </chakra.span>
                    )}
                    {line.map((token, key) => {
                      return (
                        <chakra.span
                          fontFamily="JetBrains Mono"
                          {...getTokenProps({ token, key })}
                          key={`${codeId}.${i}.${key}`}
                        />
                      );
                    })}
                  </chakra.div>
                );
              })}
            </chakra.pre>
            <chakra.div
              position="absolute"
              top={4}
              right={4}
              opacity={0}
              transition="opacity 0.2s"
              _groupHover={{
                opacity: 1,
              }}
            >
              <chakra.button
                aria-label="Copy code to clipboard"
                onClick={() => {
                  navigator.clipboard.writeText(codeString);
                  onCopy();
                }}
                color="gray.500"
                bg="white"
                _hover={{
                  color: 'gray.700',
                }}
                _dark={{
                  color: 'gray.300',
                  bg: 'blackAlpha.800',
                  _hover: {
                    color: 'white',
                  },
                }}
                fontSize="xl"
                p={2}
                rounded="md"
                boxShadow="sm"
                outline="none"
                transition="all 0.2s"
              >
                {hasCopied ? <IoMdCheckmark /> : <MdContentCopy />}
              </chakra.button>
            </chakra.div>
          </chakra.div>
        );
      }}
    </ChakraHighlight>
  );
};

const LinkedHeading = (props: HTMLChakraProps<'h2'>) => {
  const { children } = props;
  const slug = slugify(children as string, { lower: true });

  return (
    <Link
      alignItems="flex-end"
      display="flex"
      href={`#${slug}`}
      role="group"
      id={slug}
      scrollMarginTop={20}
      my={4}
    >
      <Box
        {...props}
        display="inline"
        color="gray.700"
        fontFamily="heading"
        _dark={{
          color: 'white',
        }}
      >
        {children}
      </Box>
      <chakra.span
        aria-label="anchor"
        color="blue.500"
        userSelect="none"
        fontWeight="normal"
        outline="none"
        _focus={{ opacity: 1, boxShadow: 'outline' }}
        opacity={1}
        ml="0.375rem"
        {...props}
      >
        #
      </chakra.span>
    </Link>
  );
};

const Image = (props: ImageProps) => {
  return (
    <NextImage {...props} layout="responsive" loading="lazy" quality={100} />
  );
};

const Anchor = (props: ChakraComponent<'a'>) => (
  <chakra.a target="_blank" color="blue.500" _dark={{ color: 'blue.300' }} {...props} />
);

const MDXComponents = {
  code: CodeHighlight,
  h1: (props: ChakraComponent<'h1'>) => <LinkedHeading fontSize="4xl" fontWeight="semibold" as="h1" apply="mdx.h1" {...props} />,
  h2: (props: ChakraComponent<'h2'>) => <LinkedHeading as="h2" fontSize="3xl" fontWeight="semibold" apply="mdx.h2" {...props} />,
  h3: (props: ChakraComponent<'h3'>) => <LinkedHeading as="h3" fontSize="2xl" fontWeight="semibold" apply="mdx.h3" {...props} />,
  h4: (props: ChakraComponent<'h4'>) => <LinkedHeading as="h4" fontSize="xl" fontWeight="semibold" apply="mdx.h4" {...props} />,
  hr: (props: ChakraComponent<'hr'>) => <chakra.hr apply="mdx.hr" {...props} />,
  strong: (props: BoxProps) => <Box as="strong" fontWeight="semibold" {...props} />,
  pre: Pre,
  kbd: Kbd,
  img: Image,
  br: (props: BoxProps) => <Box as="br" h={undefined} {...props} />,
  table: Table,
  th: THead,
  td: TData,
  a: Anchor,
  p: (props: ChakraComponent<'p'>) => <chakra.p apply="mdx.p" {...props} />,
  ul: (props: ChakraComponent<'ul'>) => (
    <chakra.ul px={{ base: 4, md: 0 }} py={4} apply="mdx.ul" {...props} />
  ),
  ol: (props: ChakraComponent<'ol'>) => <chakra.ol apply="mdx.ul" {...props} />,
  li: (props: ChakraComponent<'li'>) => <chakra.li pb="4px" {...props} />,
  blockquote: (props: AlertProps) => (
    <Box>
      <Alert
        as="blockquote"
        role="none"
        rounded="4px"
        status="warning"
        variant="left-accent"
        {...props}
        w="unset"
        mx={-4}
        my={4}
      />
    </Box>
  ),
} as Pick<MDXRemoteProps, 'components'>;

export default MDXComponents;
