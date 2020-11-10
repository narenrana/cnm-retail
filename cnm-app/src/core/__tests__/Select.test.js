/* eslint-disable import/no-extraneous-dependencies */
/* eslint-disable no-undef */
import React from "react";
import renderer from "react-test-renderer";
import Select from "../Select";

describe("Select test", () => {
  it("Select should match snapshot", () => {
    const component = renderer.create(<Select />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
