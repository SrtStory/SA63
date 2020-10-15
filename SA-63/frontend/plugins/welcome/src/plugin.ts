import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import statistic from './components/statistic'



export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/statistic', statistic);
 
  },
});
