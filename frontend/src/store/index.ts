/*
  Store root
*/
import Vue from 'vue'
import Vuex, { StoreOptions } from 'vuex'

import { RootState } from './types'
import { natureState } from './nature/index'
import { individualsState } from './individuals/index'
import { basePointsState } from './basePoints/index'
import { speciesState } from './species/index'
import { itemState } from './items/index'
import { attackerState } from './attacker/index'
import { weatherFieldsState } from './weatherFields/index';

import { targetSelect } from '../components/targetSelect/store/index'
import { target } from '../components/target/store/index'
import { nature } from '../components/nature/store/index'
import { speedOrder } from '../components/speedOrder/store/index'
import { moves } from '../components/moves/store/index'

Vue.use(Vuex);

const store: StoreOptions<RootState> = {
  state: {
    version: '1.0.0',
    level: 50,  // TODO:LevelStateつくる
  },
  modules: {
    target,
    targetSelect,
    nature,
    speedOrder,
    moves,
    weatherFieldsState,
    natureState,
    individualsState,
    basePointsState,
    speciesState,
    itemState,
    attackerState,
  }
}

export default new Vuex.Store<RootState>(store);
