import { useEffect, PropsWithChildren } from 'react';
import { Container, VStack } from '@chakra-ui/react';
import { useAnimation, motion } from 'framer-motion';
import { useInView } from 'react-intersection-observer';

import { Header } from '../header';
// import { Footer } from '../footer';
import { Divider } from '../divider';
import { startAnimation } from '../../utils/constants';

type Props = PropsWithChildren<{}>;

export const Layout = ({ children }: Props) => {
  const [ref, inView] = useInView();
  const controls = useAnimation();

  useEffect(() => {
    if (inView) {
      controls.start('show');
    }
  }, [controls, inView]);

  return (
    <Container
      display="flex"
      maxW="full"
      minH={{ base: 'auto', md: '100vh' }}
      centerContent
      px={0}
    >
      <VStack flex={1} spacing={10} w="full">
        <Header />
        <VStack
          spacing={10}
          flex={1}
          as="main"
          w={{ base: '90%', lg: '50%' }}
          align="flex-start"
        >
          <motion.div
            initial="hide"
            ref={ref}
            animate={controls}
            variants={startAnimation}
            style={{
              width: '100%',
            }}
          >
            {children}
          </motion.div>
        </VStack>
        <Divider />
        {/* <Footer /> */}
      </VStack>
    </Container>
  );
};
