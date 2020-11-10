/* eslint-disable import/no-extraneous-dependencies */
/* eslint-disable no-undef */
import React from "react";
import renderer from "react-test-renderer";
import Product from "../Product";

describe("Product test", () => {
  it("Product should match snapshot", () => {
    const component = renderer.create(<Product />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
