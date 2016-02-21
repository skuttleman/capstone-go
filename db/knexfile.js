// Update with your config settings.
try {
  require('dotenv').load();
} catch (err) {
  console.error(err);
}

module.exports = {
  development: {
    client: 'mysql',
    connection: 'mysql://localhost:3306/tilda'
  },
  production: {
    client: 'mysql',
    connection: process.env.DATABASE_URL
  }
};
