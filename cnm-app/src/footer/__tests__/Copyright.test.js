/* eslint-disable import/no-extraneous-dependencies */
/* eslint-disable no-undef */
import React from "react";
import renderer from "react-test-renderer";
import Copyright from "../Copyright";

describe("Copyright test", () => {
  it("Copyright should match snapshot", () => {
    const component = renderer.create(<Copyright />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
