import * as _ from 'lodash';
import store from '../../common/redux';
//import { useDispatch, useSelector } from 'react-redux';

async function   httpClient(URL,options)   {
    //const dispatch = useDispatch();
    //const { products = [], isLoading } = useSelector(state => state.commonStore);

    //store.commonStore.setLoading(true);
    return await fetch( 'http://localhost:8080'+URL,options)
      .then(async (res) => await res.json())
      .then(async res => {
        //store.commonStore.setLoading(false);
        return await res;
      })
      .catch( e => {
        //store.commonStore.setLoading(false);
        return e;
      });
  };

  export default httpClient;
 