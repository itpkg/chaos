import { moduleForComponent, test } from 'ember-qunit';
import hbs from 'htmlbars-inline-precompile';

moduleForComponent('switch-lang', 'Integration | Component | switch lang', {
  integration: true
});

test('it renders', function(assert) {
  // Set any properties with this.set('myProperty', 'value');
  // Handle any actions with this.on('myAction', function(val) { ... });

  this.render(hbs`{{switch-lang}}`);

  assert.equal(this.$().text().trim(), '');

  // Template block usage:
  this.render(hbs`
    {{#switch-lang}}
      template block text
    {{/switch-lang}}
  `);

  assert.equal(this.$().text().trim(), 'template block text');
});
