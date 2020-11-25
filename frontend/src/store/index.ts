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

import { targetSelect } from '../components/targetSelect/store/index'
import { target } from '../components/target/store/index'
import { nature } from '../components/nature/store/index'
import { speedOrder } from '../components/speedOrder/store/index'
import { moves } from '../components/moves/store/index'
import { attacker } from '../components/attacker/store/index'
import { weatherFieldsState } from './weatherFields/index';

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
    attacker,
    weatherFieldsState,
    natureState,
    individualsState,
    basePointsState,
    speciesState,
    itemState,
  }
}

export default new Vuex.Store<RootState>(store);
