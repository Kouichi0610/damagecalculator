/*
  Store root
*/
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import {counter} from './counter';
import {stats} from './stats';

// 'export default' ほかのコンポーネントでも使用できるようにする
const store = new Vuex.Store({
  modules: {
    counter: counter,
    stats: stats,
  },
  state: {
    pokename: 'ライチュウ',
  },
  getters: {
    myname: state => {
      return state.pokename + store.getters['counter/count'];
    },
    statsString: (/*state*/) => {
      return store.getters['stats/toString'];
    },


  },
})

console.log('myname:::' + store.getters.myname);


export default store
