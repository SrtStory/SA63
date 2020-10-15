import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';



import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="" text="Home" />
    <SidebarItem icon={CreateComponentIcon} to="statistic" text="statistic" />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" />
   
   
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarUserSettings  />
    <SidebarPinButton />
  </Sidebar>
);