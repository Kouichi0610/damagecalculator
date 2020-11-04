import { GetterTree } from 'vuex';
import { DefaultSpeedCorrector, SpeedOrderState, SpeedCorrector } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<SpeedOrderState, RootState> = {
  hasList: (state: SpeedOrderState) => {
    return state.list.length > 0;
  },
  list: (state: SpeedOrderState) => {
    return state.list;
  },
  abilityOwnerCorrector: (state: SpeedOrderState): SpeedCorrector => {
    return state.abilityOwner;
  },
  abilityOtherCorrector: (state: SpeedOrderState): SpeedCorrector => {
    return state.abilityOther;
  },
  itemCorrectors: (state: SpeedOrderState): SpeedCorrector[] => {
    return state.item;
  },
}
