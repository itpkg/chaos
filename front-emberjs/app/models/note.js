import Model from 'ember-data/model';
import attr from 'ember-data/attr';

export default Model.extend({
  id: attr(),
  title: attr(),
  body: attr(),
  updatedAt: attr()
});
