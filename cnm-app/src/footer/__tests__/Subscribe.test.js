/* eslint-disable import/no-extraneous-dependencies */
/* eslint-disable no-undef */
import React from "react";
import renderer from "react-test-renderer";
import Subscribe from "../Subscribe";

describe("Subscribe test", () => {
  it("Subscribe should match snapshot", () => {
    const component = renderer.create(<Subscribe />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
