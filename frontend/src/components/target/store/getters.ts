import { GetterTree } from 'vuex';
import { TargetState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<TargetState, RootState> = {
  species: (state: TargetState) => {
    return state.species;
  }
}
