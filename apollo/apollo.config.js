// Apollo Studio configuration file
module.exports = {
  client: {
    service: {
      name: "prototype01",
      url: "http://localhost:8080/graphql",
    },
    includes: ["./queries/**/*.{js,ts,graphql}"],
  },
};
