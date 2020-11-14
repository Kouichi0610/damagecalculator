/*
  Store root
*/
import Vue from 'vue'
import Vuex, { StoreOptions } from 'vuex'

import { RootState } from './types'
import { targetSelect } from '../components/targetSelect/store/index'
import { target } from '../components/target/store/index'
import { nature } from '../components/nature/store/index'
import { speedOrder } from '../components/speedOrder/store/index'
import { attacker } from '../components/attacker/store/index'
import { moves } from '../components/moves/store/index'

Vue.use(Vuex);

const store: StoreOptions<RootState> = {
  state: {
    version: '1.0.0',
    level: 50,
  },
  modules: {
    target,
    targetSelect,
    nature,
    speedOrder,
    attacker,
    moves,
  }
}

export default new Vuex.Store<RootState>(store);
