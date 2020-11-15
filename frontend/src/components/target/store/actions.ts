import { ActionTree } from 'vuex';
import { RootState } from '@/store/types'
import { TargetState } from './types'
import { StatePatternsLoader } from './statePattern'
import { SpeciesLoader } from './species'

export const actions: ActionTree<TargetState, RootState> = {
  setCurrentAbility: ({commit}, name: string) => {
    commit('setCurrentAbility', name);
  },
  getStatsPattern: ({commit}, args: StatePatternsLoader) => {
    args.load()
    .then((patterns) => {
      commit('setStatsPattern', patterns);
    });
  },
  getSpecies: ({commit}, loader: SpeciesLoader): Promise<boolean> => {
    return new Promise((resolve, reject) => {
      loader.load()
      .then((species) => {
        commit('setSpecies', species);
        resolve(true);
      })
      .catch((e) => {
        reject(false);
      });
    });
  },
}

