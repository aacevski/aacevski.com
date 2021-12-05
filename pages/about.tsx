import {
  Text,
  Heading,
  Box,
  Tab,
  TabProps,
  Tabs,
  TabList,
  TabPanels,
  TabPanel,
  Divider,
  VStack,
  Link,
  useColorModeValue as mode,
} from '@chakra-ui/react';

import { TopTracks } from '../src/components/top-tracks';

const CustomTab = ({ children }: TabProps) => {
  return (
    <Tab
      opacity={0.5}
      fontWeight="bold"
      fontSize="xl"
      _selected={{
        fontSize: '2xl',
        opacity: 1,
      }}
      _focus={{
        outline: 'none',
      }}
    >
      {children}
    </Tab>
  );
};

const About = () => {
  return (
    <VStack>
      <Tabs variant="unstyled">
        <TabList h={12} display="flex" justifyContent="center">
          <CustomTab>Story</CustomTab>
          <CustomTab>Music</CustomTab>
        </TabList>
        <TabPanels>
          <TabPanel>
            <VStack spacing={10}>
              <VStack align="flex-start">
                <Text fontSize="3rem" mb={2}>
                  👋
                </Text>
                <Heading size="lg">
                  Hey, I’m Andrej! I’m a{' '}
                  <Box
                    as="span"
                    bgGradient="linear(to-l, #12c2e9, #c471ed, #f64f59)"
                    bgClip="text"
                    textDecoration="underline"
                  >
                    web developer
                  </Box>{' '}
                  located in Macedonia, other than that I love photography and
                  long car rides.
                </Heading>
              </VStack>

              <VStack>
                <Heading
                  as="h3"
                  fontSize="lg"
                  color="gray"
                  fontWeight="medium"
                  textAlign="justify"
                >
                  With over a year of experience in front end development, I’m
                  currently working as a web developer, with an amazing team of
                  great developers. I’ve experience with React and Node.js.
                </Heading>
                <Heading
                  as="h3"
                  fontSize="lg"
                  color="gray"
                  fontWeight="medium"
                  mt={4}
                  textAlign="justify"
                >
                  Outside of work, I’m a student at the{' '}
                  <Link
                    href="https://www.finki.ukim.mk/en"
                    color={mode('black', 'white')}
                  >
                    Faculty of Computer Science and Engineering
                  </Link>{' '}
                  and I’m a member of the student organization{' '}
                  <Link
                    color={mode('black', 'white')}
                    href="https://eestec.mk/"
                  >
                    EESTEC
                  </Link>
                  . I’ve attented a number of events organized by EESTEC and the
                  faculty I’m attending, regarding various technologies.
                </Heading>
                <Heading
                  as="h3"
                  fontSize="lg"
                  color="gray"
                  fontWeight="medium"
                  mt={4}
                  textAlign="justify"
                >
                  In my free time, I love to take photos of amazing views and
                  visit new places. On top of all that, like every tech nerd, I
                  love to learn more about technology - even outside of my
                  field. The last thing I’ve tried is Arduino and it’s been
                  super fun, in my opinion you can never know too much.
                </Heading>
              </VStack>
            </VStack>
          </TabPanel>
          <TabPanel>
            <VStack spacing={4} w="full">
              <Heading fontSize="lg" color="#7e7e7e" fontWeight="medium">
                Welcome to my music corner, peak around and maybe you’ll find
                your new favorite song!
              </Heading>
              <Divider />
              <Heading fontSize="2xl">Top Tracks</Heading>
              <VStack align="flex-start" w="full">
                <TopTracks />
              </VStack>
            </VStack>
          </TabPanel>
        </TabPanels>
      </Tabs>
    </VStack>
  );
};

export default About;
