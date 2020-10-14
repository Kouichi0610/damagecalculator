/*
  Store root
*/
import Vue from 'vue'
import Vuex, { StoreOptions } from 'vuex'

import { RootState } from './types'
import { profile } from './profile/index'

Vue.use(Vuex);

const store: StoreOptions<RootState> = {
  state: {
    version: '1.0.0'
  },
  modules: {
    profile,
  }
}

export default new Vuex.Store<RootState>(store);

/*
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
    statsString: () => {
      return store.getters['stats/toString'];
    },


  },
})

console.log('myname:::' + store.getters.myname);

export default store
*/