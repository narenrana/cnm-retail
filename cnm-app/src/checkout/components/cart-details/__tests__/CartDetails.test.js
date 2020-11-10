/* eslint-disable import/no-extraneous-dependencies */
/* eslint-disable no-undef */
import React from "react";
import renderer from "react-test-renderer";
import CartDetails from "../CartDetails";

describe("CartDetails test", () => {
  it("CartDetails should match snapshot", () => {
    const component = renderer.create(<CartDetails />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
