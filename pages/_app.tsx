import React from 'react';
import { AppProps } from 'next/app';
import { ChakraProvider } from '@chakra-ui/react';
import { Global } from '@emotion/react';
import { NextSeo } from 'next-seo';

import { Layout } from '../src/components/layout';
import theme from '../theme/index';
import { Fonts } from '../theme/fonts';
import '../styles/globals.css';

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ChakraProvider theme={theme}>
      <NextSeo
        title="Andrej Acevski"
        description="Web developer who spends too much time coding."
        twitter={{
          cardType: 'summary_large_image',
          handle: '@aacevski',
        }}
        openGraph={{
          url: 'https://aacevski.com',
          title: 'Andrej Acevski',
          description: 'Web developer who spends too much time coding.',
          locale: 'en_US',
          images: [
            {
              url: 'https://i.imgur.com/YaQt1UM.jpg',
              width: 1200,
              height: 630,
              alt: 'Andrej Acevski',
              type: 'image/jpeg',
            },
          ],
        }}
      />
      <Global styles={Fonts} />
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </ChakraProvider>
  );
}

export default MyApp;
