/* eslint-disable import/no-extraneous-dependencies */
/* eslint-disable no-undef */
import React from "react";
import renderer from "react-test-renderer";
import Checkout from "../Checkout";

describe("Checkout test", () => {
  it("Checkout should match snapshot", () => {
    const component = renderer.create(<Checkout />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
