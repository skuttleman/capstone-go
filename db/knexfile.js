// Update with your config settings.
try {
  var parent = __dirname.split('/');
  parent.pop();
  parent = parent.join('/');
  require('dotenv').config({ path: parent + '/.env' });
} catch (err) {
  console.log(err);
}


module.exports = {
  development: {
    client: 'mysql',
    connection: {
      database: 'tilda',
      user:  'root',
      host: 'localhost',
      port: 3306
    }
  },
  production: {
    client: 'mysql',
    connection: process.env.DATABASE_URL
  }
};
