/* eslint-disable import/no-extraneous-dependencies */
/* eslint-disable no-undef */
import React from "react";
import renderer from "react-test-renderer";
import ProductHome from "../ProductHome";

describe("ProductHome test", () => {
  it("ProductHome should match snapshot", () => {
    const component = renderer.create(<ProductHome />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
