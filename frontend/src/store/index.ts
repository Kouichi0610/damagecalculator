/*
  Store root
*/
import Vue from 'vue'
import Vuex, { StoreOptions } from 'vuex'

import { RootState } from './types'
import { targets } from './targets/index'
import { speedList } from './speedlist/index'
import { targetSelect } from '../components/targetSelect/store/index'
import { target } from '../components/target/store/index'
import { nature } from "../components/nature/store/index"

Vue.use(Vuex);

const store: StoreOptions<RootState> = {
  state: {
    version: '1.0.0',
    level: 50,
  },
  modules: {
    targets,
    speedList,
    targetSelect,
    target,
    nature,
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