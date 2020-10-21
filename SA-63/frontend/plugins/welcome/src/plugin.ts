import { createPlugin } from '@backstage/core';

import statistic from './components/statistic'
import Login from './components/Login'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/statistic', statistic);
 
  },
});
