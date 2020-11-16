import axios from 'axios'

/*
  種族固有の
*/
export class Species {
  readonly name: string = '';
  readonly types: string = '';
  readonly weight: number = 0;
  readonly species: number[] = [0,0,0,0,0,0];
  readonly abilities: string[] = [];

  hasTarget(): boolean {
    return this.name.length > 0;
  }
  hp(): number {return this.species[0];}
  attack(): number {return this.species[1];}
  defense(): number {return this.species[2];}
  spAttack(): number {return this.species[3];}
  spDefense(): number {return this.species[4];}
  speed(): number {return this.species[5];}

  static default(): Species {
    return new Species(null);
  }
  constructor(data: any) {
    if (data == null) {
      return;
    }
    this.name = data.Name;
    this.types = data.Types;
    this.weight = data.Weight;
    this.species = data.Species;
    this.abilities = data.Abilities;
  }
}

/*
  Species取得
*/
export class SpeciesLoader {
  private name: string;

  constructor(name: string) {
    this.name = name;
  }

  public load(): Promise<Species> {
    return new Promise((resolve, reject) => {
      axios.get('get_species', {
        params: {
          Name: this.name,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let sp = JSON.parse(json);
        resolve(new Species(sp));
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject(new Species(null));
      });
    });
  }
}
