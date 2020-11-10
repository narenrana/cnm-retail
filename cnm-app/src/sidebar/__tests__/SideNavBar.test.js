
// Auto-generated do not edit


/* eslint-disable import/no-extraneous-dependencies */
/* eslint-disable no-undef */
import React from 'react';
import renderer from 'react-test-renderer';
import SideNavBar from '../SideNavBar';


describe('SideNavBar test', () => {
  it('SideNavBar should match snapshot', () => {
    const component = renderer.create(<SideNavBar
       />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
