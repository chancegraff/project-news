var express = require("express");
var cors = require("cors");
var graphqlHTTP = require("express-graphql");
var gql = require("graphql-tag");
var { buildASTSchema } = require("graphql");

const ARTICLES = [
  { title: "John Doe", address: "google.com" },
  { title: "Jane Doe", address: "google.com" },
];

const schema = buildASTSchema(gql`
  type Query {
    articles: [Article]
    article(id: ID!): Article
  }

  type Article {
    id: ID
    title: String
    address: String
  }
`);

const mapArticle = (post, id) => post && ({ id, ...post });

const root = {
  articles: () => ARTICLES.map(mapArticle),
  article: ({ id }) => mapArticle(ARTICLES[id], id),
};

const app = express();
app.use(cors());
app.use('/graphql', graphqlHTTP({
  schema,
  rootValue: root,
  graphiql: true,
}));

const port = process.env.PORT || 4000
app.listen(port);
console.log(`Running a GraphQL API server at localhost:${port}/graphql`);
