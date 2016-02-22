
exports.up = function(knex, Promise) {
  return knex.schema.createTable('players', function(table) {
    table.increments('id');
    table.string('social_id');
    table.string('image');
    table.string('email');
    table.string('name');
  }).then(function() {
    return knex.schema.createTable('games', function(table) {
      table.increments('id');
      table.integer('player1_id').unsigned().references('id').inTable('players').onDelete('CASCADE');
      table.integer('player2_id').unsigned().references('id').inTable('players').onDelete('CASCADE');
      table.string('game_state_id');
    });
  });
};

exports.down = function(knex, Promise) {
  return knex.schema.dropTableIfExists('games').then(function() {
    return knex.schema.dropTableIfExists('players');
  });
};
