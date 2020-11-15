import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types'
import { TargetState, Species } from './types'
import { StatePatternsLoader } from './statePattern'

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
  getSpecies: ({commit}, name: string): Promise<string> => {
    return new Promise((resolve, reject) => {
      axios.get('get_species', {
        params: {
          Name: name,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let sp = JSON.parse(json);
        commit('setName', sp.Name);
        commit('setTypes', sp.Types);
        commit('setWeight', sp.Weight);
        let species: Species = {hp: sp.Species[0], at: sp.Species[1], df: sp.Species[2], sa: sp.Species[3], sd: sp.Species[4], sp: sp.Species[5]};
        commit('setSpecies', species);
        commit('setAbilities', sp.Abilities);
        resolve(sp.Name);
      })
      .catch((e) => {
          console.log('failed:' + e);
          reject('');
      });
    });
  },
}

