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
import { weatherFieldsState } from './weatherFields/index';
import { attackerState } from './attacker/index'
import { speedState } from './speed/index'
import { targetSelectState } from './targetSelect/index'

Vue.use(Vuex);

const store: StoreOptions<RootState> = {
  state: {
    version: '1.0.0',
    level: 50,  // TODO:LevelStateつくる
  },
  modules: {
    targetSelectState,
    weatherFieldsState,
    natureState,
    individualsState,
    basePointsState,
    speciesState,
    itemState,
    attackerState,
    speedState,
  }
}

export default new Vuex.Store<RootState>(store);
