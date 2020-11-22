import { GetterTree } from 'vuex';
import { Individuals, IndividualsState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<IndividualsState, RootState> = {
  individuals: (state: IndividualsState): Individuals => {
    return state.targetsIndividuals;
  },
}
