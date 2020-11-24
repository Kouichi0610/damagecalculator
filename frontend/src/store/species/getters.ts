import { GetterTree } from 'vuex';
import { SpeciesState, PokeData } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<SpeciesState, RootState> = {
  pokeData: (state: SpeciesState): PokeData => {
    return state.data;
  }
}
