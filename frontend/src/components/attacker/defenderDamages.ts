import axios from 'axios';

export class DefenderDamages {
  /*
  
  */
  getSample() {
    axios.get('defender_damages', {
      params: {
        Level: 50,
        BaseHP: 0,
        BaseAttack: 252,
        BaseDefense: 0,
        BaseSpAttack: 0,
        BaseSpDefense: 0,
        BaseSpeed: 252,
            Individuals: 'Max',
        Name: 'ピカチュウ',
        Move: 'かみなりパンチ',
        Ability: 'せいでんき',
        Nature: 'いじっぱり',
        Item: '',
        Condition: '',
      }
    })
    .then((response) => {
      let json = JSON.stringify(response.data);
      //let descriptions = JSON.parse(json);
      console.log('' + json);
      console.log('success.');
      return true;
    })
    .catch((e) => {
      console.log('failed:' + e);
      return false;
    });
  }


}
