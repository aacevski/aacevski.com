import React from 'react';
import { Text, VStack } from '@chakra-ui/react';

export const Footer = () => {
  return (
    <VStack w="full">
      <Text fontSize="sm" color="gray" mb={10}>
        {new Date().getFullYear()} © Andrej Acevski — Keep on rockin’ ⚡
      </Text>
    </VStack>
  );
};
