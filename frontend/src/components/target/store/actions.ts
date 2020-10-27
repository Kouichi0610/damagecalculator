import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types'
import { TargetState, Species } from './types'

export const actions: ActionTree<TargetState, RootState> = {
  getSpecies: ({commit}, name: string) => {
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
      return true;
    })
    .catch((e) => {
        console.log('failed:' + e);
        return false;
    });
  }
}

