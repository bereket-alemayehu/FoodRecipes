import { defineNuxtPlugin } from "#app";
import {
  ApolloClient,
  createHttpLink,
  InMemoryCache,
} from "@apollo/client/core";
import { provideApolloClient } from "@vue/apollo-composable";

export default defineNuxtPlugin((nuxtApp) => {
  const httpLink = createHttpLink({
    uri: "YOUR_HASURA_GRAPHQL_ENDPOINT", // Replace with your Hasura endpoint
    headers: {
      authorization: `Bearer ${localStorage.getItem("token")}`, // Replace with your token storage
    },
  });

  const cache = new InMemoryCache();

  const apolloClient = new ApolloClient({
    link: httpLink,
    cache,
  });

  provideApolloClient(apolloClient);
});
