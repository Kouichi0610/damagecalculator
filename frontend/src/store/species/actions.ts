import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { SpeciesState, PokeDataLoader } from './types'

export const actions: ActionTree<SpeciesState, RootState> = {
  load: ({ commit }, name: string) => {
    let loader = new PokeDataLoader();
    loader.load(name)
    .then((d) => {
      commit('setData', d);
    });
  },
}
